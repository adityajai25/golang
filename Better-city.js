const axios = require("axios")
const city1 = "Srinagar";
const city2 = "Pune";
const condition = "hot";
const getWeather = async(city) => {
    const url = `https://quest.squadcast.tech/api/RA2011003020090/weather/get?q=${city}`;
    const response = await axios.get(url);
    return response.data;
}
const compareWeather = async() => {
    const city1Weather = await getWeather(city1);
    const city2Weather = await getWeather(city2);
    if (condition === "hot") {
        const city1Param = city1Weather.main.temp;
        const city2Param = city2Weather.main.temp;
        if (city1Param > city2Param) {
            console.log(`${city1} is better`);
        } else {
            console.log(`${city2} is better`);
        }
    } else if (condition === "cold") {
        const city1Param = city1Weather.main.temp;
        const city2Param = city2Weather.main.temp;
        if (city1Param < city2Param) {
            console.log(`${city1} is better`);
        } else {
            console.log(`${city2} is better`);
        }
    } else if (condition === "windy") {
        const city1Param = city1Weather.wind.speed;
        const city2Param = city2Weather.wind.speed;
        if (city1Param > city2Param) {
            console.log(`${city1} is better`);
        } else {
            console.log(`${city2} is better`);
        }
    } else if (condition === "cloudy") {
        const city1Param = city1Weather.clouds.all;
        const city2Param = city2Weather.clouds.all;
        if (city1Param > city2Param) {
            console.log(`${city1} is better`);
        } else {
            console.log(`${city2} is better`);
        }
    } else if (condition === "rainy") {
        const city1Param = city1Weather.rain ? city1Weather.rain["1h"] : 0;
        const city2Param = city2Weather.rain ? city2Weather.rain["1h"] : 0;
        if (city1Param > city2Param) {
            console.log(`${city1} is better`);
        } else {
            console.log(`${city2} is better`);
        }
    } else {
        console.log("Invalid condition");
    }
}
compareWeather();