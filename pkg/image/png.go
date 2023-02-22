package image

import (
	"encoding/binary"
	"io"
)

type PngChunk struct {
	Length int
	Name   string
	Data   []byte
	CRC    uint32
}

func DecodePngTextChunk(r io.Reader) ([]PngChunk, error) {
	d := decoder{r: r, chunks: make([]PngChunk, 0)}
	return d.chunks, d.parse("tEXt")
}

type decoder struct {
	r      io.Reader
	chunks []PngChunk
}

func (d *decoder) parse(targetChunk string) error {
	sig := make([]byte, 8)
	if _, err := d.r.Read(sig); err != nil {
		return err
	}

	for {
		err := d.parseChunk()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		if d.chunks[len(d.chunks)-1].Name == targetChunk {
			return nil
		}
	}
}

func (d *decoder) parseChunk() error {
	head := make([]byte, 8)
	if _, err := d.r.Read(head); err != nil {
		return err
	}

	chunkLen := binary.BigEndian.Uint32(head[:4])
	chunkName := string(head[4:8])
	chunkData := make([]byte, chunkLen)

	if _, err := d.r.Read(chunkData); err != nil {
		return err
	}

	crc := make([]byte, 4)
	if _, err := d.r.Read(crc); err != nil {
		return err
	}

	d.chunks = append(d.chunks, PngChunk{
		Length: int(chunkLen),
		Name:   chunkName,
		Data:   chunkData,
		CRC:    binary.BigEndian.Uint32(crc),
	})
	return nil
}
