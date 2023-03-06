import { Card, Grid, ImageList, ImageListItem, Typography } from "@mui/material"
import { Image } from "../../../Models";
import './ImagesList.css';

export interface ImagesListProps {
    images: Image[];
}

export const ImagesList = (props: ImagesListProps) => {
    const renderImage = (image: Image) => (
        <ImageListItem key={image.name}>
            <img className='images_list__img' src={image.url} loading='lazy'/>
        </ImageListItem>
    );

    return (
        <Grid container spacing={0} rowSpacing={0}>
            {props.images.map((image) => <Grid item md={3} sm={4} key={image.pictureId}>
                <img className='images_list__img'
                    src={image.url}
                    loading='lazy'
                    width={image.width}
                    height={image.height}
                />
            </Grid>)}
        </Grid>
    );

    return (
        <ImageList cols={3} variant='quilted'>
            {props.images.map((image) => renderImage(image))}
        </ImageList>
    );  
}
