export type Image = {
  pictureId: string,
  name: string,
  url: string,
  width: number,
  height: number
};

export type ImageList = {
  items: Image[],
  count: number,
  page: number
}