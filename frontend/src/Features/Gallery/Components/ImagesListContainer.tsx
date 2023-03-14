import { Container, Loading } from "@nextui-org/react";
import { useEffect, useState } from "react";
import { getImageFeed } from "../../../Api/Feed";
import { Image } from '../../../Models';
import { useInfiniteScroll } from "../../../Shared/Hooks";
import { ImagesList } from "./ImagesList"

const pageSize = 24;

export const ImagesListContainer = () => {
  const [imagesPages, setImagesPages] = useState<Image[][]>([]);
  const [isLoading, setIsLoading] = useInfiniteScroll(loadMoreImages);

  function loadMoreImages() {
    let pageIndex = imagesPages.length;

    getImageFeed(pageIndex, pageSize)
      .then(imagesList => {
        setImagesPages(pervState => ([...pervState, imagesList.items]));

        if (imagesList.items.length === pageSize)
          setIsLoading(false);
      });
  };

  return <>
    {imagesPages.map((page, pageIdx) => <ImagesList images={page} key={pageIdx} />)}
    <Container css={{dflex: 'center', p: 48}}>
      <Loading />
    </Container>
  </>;
}