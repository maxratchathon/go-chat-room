import type { AppProps } from 'next/app'
import { WebSocketProvider } from '../src/components/WebSocketProvider'
import '../src/styles/globals.css'
import Head from 'next/head'

function MyApp({ Component, pageProps }: AppProps) {
  return (
    <WebSocketProvider>
      <Head>
        <title>Go Chat Room</title>
        <meta name="description" content="Real-time chat application" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <Component {...pageProps} />
    </WebSocketProvider>
  )
}

export default MyApp