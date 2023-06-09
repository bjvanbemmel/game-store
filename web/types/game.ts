import Developer from '~/types/developer'
import Genre from '~/types/genre'
import Media from '~/types/media'

export default interface Game {
    id: number,
    title: string,
    description: string,
    thumbnail: string,
    price: number,
    media: Array<Media>,
    genres: Array<Genre>,
    developers: Array<Developer>,
}
