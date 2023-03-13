import { Grid } from "@nextui-org/react";
import { Image } from "../../../Models";
import './ImagesList.css';

export interface ImagesListProps {
    images: Image[];
}

export const ImagesList = (props: ImagesListProps) => {
    return (
        <Grid.Container gap={0}>
            {props.images.map((image) => <Grid md={3} sm={4} key={image.pictureId}>
                <img className='images_list__img'
                    src={image.url}
                    loading='lazy'
                    width={image.width}
                    height={image.height}
                />
            </Grid>)}
        </Grid.Container>
    );
}
