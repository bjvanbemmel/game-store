import Media from "~/types/media";
import MediaType from "~/types/media_type";

export default class YouTube implements Media {
    type: MediaType
    uri: string

    constructor(uri: string) {
        this.type = MediaType.YouTube

        if (! /^https\:\/\/youtube\.com\/embed\//.test(uri)) {
            throw new Error("Given URI does not start with 'https://youtube.com/embed'")
        }
        this.uri = uri
    }
}
