import Developer from '~/types/developer'
import Genre from '~/types/genre'
import Media from '~/types/media'
import PlayerCount from '~/types/player_count'

export default interface Game {
    id: number,
    title: string,
    description: string,
    thumbnail: string,
    price: number,
    media: Array<Media>,
    genres: Array<Genre>,
    developers: Array<Developer>,
    release_date: Date,
    rating: number,
    player_count: Array<PlayerCount>,
}
