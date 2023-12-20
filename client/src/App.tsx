import './App.css'

import { useEffect, useState } from 'react'
import useWebSocket, { ReadyState } from "react-use-websocket"
import { SendJsonMessage } from 'react-use-websocket/dist/lib/types'

type notification = {
    orders: number,
    sales: number,
    products: number,
}

function App() {

    const WS_URL = "ws://localhost:8080/ws"

    const [notifications, setNotifications] = useState<notification>()

    const { sendJsonMessage, lastJsonMessage, readyState }: { sendJsonMessage: SendJsonMessage, lastJsonMessage: notification, readyState: ReadyState } = useWebSocket(
        WS_URL, {
        share: false,
        shouldReconnect: () => true,
    })

    useEffect(() => {
        console.log("Connection state changed")
        if (readyState === ReadyState.OPEN) {
            sendJsonMessage({
                event: "subscribe",
                data: [
                    "orders",
                    "sales",
                    "products"
                ],
            })
        }
    }, [readyState])

    useEffect(() => {
        console.log("Last message changed")
        if (lastJsonMessage) {
            console.log(lastJsonMessage)
            setNotifications(lastJsonMessage)
        }
    }, [lastJsonMessage])

    return (
        <>
            <h1>Notifications</h1>
            <div className="container">
                <div className="card">
                    <h2>Orders: {notifications?.orders}</h2>
                </div>
                <div className="card">
                    <h2>Sales: {notifications?.sales}</h2>
                </div>
                <div className="card">
                    <h2>Products: {notifications?.products}</h2>
                </div>
            </div>
        </>
    )
}

export default App

// ------------------------------------------------------------------------------------------------------------------------------------------------------------------------

// const { lastJsonMessage } = useWebSocket(
//     WS_URL, {
//     share: false,
//     shouldReconnect: () => true,
// })

// // Run when the connection state (readyState) changes
// useEffect(() => {
//     console.log("Connection state changed")
//     if (readyState === ReadyState.OPEN) {
//         sendJsonMessage({
//             event: "subscribe",
//             data: "TEST",
//         })
//     }
// }, [readyState])