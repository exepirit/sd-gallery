import { useEffect, useState } from "react";
import { getImageFeed } from "../../../Api/Feed";
import { Image } from '../../../Models';
import { useInfiniteScroll } from "../../../Shared/Hooks";
import { ImagesList } from "./ImagesList"

const pageSize = 24;

export const ImagesListContainer = () => {
  const [images, setImages] = useState<Image[]>([]);
  const [isLoading, setIsLoading] = useInfiniteScroll(loadMoreImages);

  function loadMoreImages() {
    let pageIndex = Math.floor(images.length / pageSize);

    getImageFeed(pageIndex, pageSize)
      .then(imagesList => {
        setImages(pervState => ([...pervState, ...imagesList.items]));

        if (imagesList.items.length === pageSize)
          setIsLoading(false);
      });
  };

  return <ImagesList images={images} />;
}