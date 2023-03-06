import axios from "axios";
import { ImageList } from "../Models";

export const getImageFeed = (page: number, count: number): Promise<ImageList> => 
    axios.get<ImageList>(`/api/feed?page=${page}&count=${count}`)
        .then(res => res.data);