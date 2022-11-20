import axios from "axios"

const token = localStorage.getItem("JWT") || "";

export const AxiosInstance = axios.create({
    baseURL: "http://localhost:8000/",
    headers: {
        Authorization: `Bearer ${token}`,
        Accept: 'application/json',
        'Content-Type': 'application/json',
    },
});