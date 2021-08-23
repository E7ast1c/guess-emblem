import axios from './axios.js'

export const getEmblem = async () => {
    // try {
    //
    //     console.log(response);
    // } catch (error) {
    //     console.error(error);
    // }
    return (await axios.get('get-random')).data;
}

