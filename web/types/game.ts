import Developer from 'types/developer'
import Genre from 'types/genre'

export default interface Game {
    id: number,
    title: string,
    description: string,
    thumbnail: string,
    price: number,
    genres: Array<Genre>,
    developers: Array<Developer>,
}
