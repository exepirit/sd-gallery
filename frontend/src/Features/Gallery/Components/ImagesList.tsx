import { Grid } from "@nextui-org/react";
import { Link } from "react-router-dom";
import { Image } from "../../../Models";
import './ImagesList.css';

export interface ImagesListProps {
    images: Image[];
}

export const ImagesList = (props: ImagesListProps) => {
    return (
        <Grid.Container gap={0}>
            {props.images.map((image) => <Grid md={3} sm={4} key={image.pictureId}>
                <Link to={`/p/${image.pictureId}`}>
                    <img className='images_list__img'
                        src={image.url}
                        loading='lazy'
                        width={image.width}
                        height={image.height}
                    />
                </Link>
            </Grid>)}
        </Grid.Container>
    );
}
