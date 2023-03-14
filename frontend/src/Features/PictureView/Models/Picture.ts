export type PictureSize = {
    width: number,
    height: number
};

export type PictureGenerateInfo = {
    prompt: string,
    negativePrompt: string,
    steps: number,
    size: string,
    seed: string,
    sampler: string,
    cfgScale: number,
    modelName: string,
    modelHash: string
};

export type Picture = {
    id: string,
    name: string,
    size: PictureSize,
    tags: string[],
    info: PictureGenerateInfo
};