import axios from './axios.js'

export const getEmblem = async () => {
    return (await axios.get('get-random')).data;
}

