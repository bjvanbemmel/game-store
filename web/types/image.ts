import Media from "~/types/media";
import MediaType from "~/types/media_type";

export default class Image implements Media {
    type: MediaType
    uri: string

    constructor(uri: string) {
        this.type = MediaType.Image
        this.uri = uri
    }
}
