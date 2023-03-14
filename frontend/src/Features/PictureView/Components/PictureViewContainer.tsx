import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import { getPicture } from "../Api";
import { Picture } from "../Models";
import { PictureView } from "./PictureView";

export const PictureViewContainer = () => {
  const [picture, setPicture] = useState<Picture>();
  const {pictureId} = useParams();

  useEffect(() => {
    getPicture(pictureId!)
        .then(picture => setPicture(picture))
        .catch(() => {});
  }, []);
  
  return <>
    {picture && <PictureView picture={picture} />}
  </>;
};