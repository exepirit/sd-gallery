import axios from "axios";
import { Picture } from "../Models";

export const getPicture = (id: string): Promise<Picture> => axios
    .get(`/api/pictures/${id}`)
    .then(res => res.data);