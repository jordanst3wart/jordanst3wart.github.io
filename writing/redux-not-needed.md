---
title: Redux is not needed anymore
date: '2024-02-28'
tags: ['redux', 'frontend', 'react.js', 'state management', 'javascript', 'typescript']
summary: Redux is not needed anymore, there is better alternatives, and a lot of frontends don't need to manage state anymore.
---

I don't think Redux is needed anymore. And I think this is because of three different reasons:

1. A lot of frontends should not have a lot of state in them
2. Some frontends use redux to check if elements are loaded which is just a bad practice, and react query is better for this
3. And finally if you do have complicated state in the frontend there are better, and simplier alternatives, such as Zustand, and Jotai, react.js useState, preact useSignal

### 1. A lot of frontends should not have a lot of state in them

I want you to have a look at a stream on twitch.tv, and see how much state is in the frontend. Twitch.tv has an extremely interactive frontend, and should be using something to manage state in the frontend. It has "local machine state", or "browser state" such as a live audio feed, and live video feed. It also has a lot of user information that needs to be passed down to the frontend such as a live message feed. It has local browser state, like is the microphone streaming to the browser, and complex application information like a live chat. This makes it very strategic to carefully manage state in the frontend.

If your application has to carefully manage frontend state like twitch.tv then it is a great idea to have state in the frontend. Otherwise, you should central state more from the server. It is extremely rare to have require such complicated state in the frontend. There is only a few categories that should require if like online games, application that need to be offline, video streaming, and live feed websites. Other websites if probably better just to store state in the server.

Separately, a lot of frontends are moving from the traditions single page app (or SPA) model to be rendered on the server (or use server side rendering). This is so the application can render faster, and have a better user experience. This is becoming more popular with Next.js, and react.

### 2. Some frontends use redux to check if elements are loaded which is just a bad practice

In some React, or redux codebases you have:

```javascript
const dispatch = useDispatch()
const { data, isLoading, error } = useSelector((state) => state.data)
// potentially some use useEffect
if (isLoading) return <div>Loading...</div>
```

The `useDispatch` and `useSelector` are from the `react-redux` library. The `useDispatch` is used to dispatch an action to the redux store, and the `useSelector` is used to get the state from the redux store. The `useEffect` is used to check if the data is loading, and if it is loading then it will return a loading message.

This requires a store to be created, and a reducer to be created. The reducer is created like:

```javascript
function dataReducer(state = initialState, action) {
  switch (action.type) {
    case 'FETCH_DATA_BEGIN':
      return {
        ...state,
        isLoading: true,
        error: null,
      }
    case 'FETCH_DATA_SUCCESS':
      return {
        ...state,
        isLoading: false,
        data: action.payload.data,
      }
    default:
      return state
  }
}

const rootReducer = combineReducers({
  data: dataReducer,
})
export const store = createStore(rootReducer)
```

This is extremely heavily for the purpose of checking if the data is loading. You are creating a state machine to handle web requests. If the website is rendered on the server as well, this logic is not required. But if it is rendered on the client it's simpler to use `useState`, and better to use `react-query` to check if the data is loading.

You can use `useState`, like:

```javascript
const [data, setData] = useState([])
const [isLoading, setIsLoading] = useState(false)

useEffect(async () => {
  setIsLoading(true)
  const data = await fetch('/some-api-endpoint')
  setData(data)
  setIsLoading(false)
}, [])

if (isLoading) return <div>Loading...</div>
```

This is a lot simpler than introducing a state container, and handling the state in the component and the reducer.

This works, but doesn't look the best. This is where `react-query` comes in. `react-query` is a library that is used to handle web requests, and check if the data is loading. It is a lot simpler than using `useState`, and `useEffect` to check if the data is loading. You can use `react-query` like:

```javascript
const { data, isLoading, error } = useQuery('data', () => fetch('/some-api-endpoint'))
if (isLoading) return <div>Loading...</div>
```

This looks a lot nicer than useEffect. You do have to learn the react-query library though. It also might not be the best managing state that says on the frontend, like writing cookies, saving to local storage, is the camera on, or off and is the microphone on, or off. It's worth mentioning a lot of libraries are changing to `useSignal` like Solid.js, preact, and qwik. You can use preact signals in react apps by using `@preact/signals-react` library.

The `useSignal` is using the observer pattern, rather than the state machine pattern used by redux.

As an aside, apparently useContext might not be the best, even though it's available from the react library.

### 3. And finally if you do have complicated state in the frontend there are better, and simplier alternatives, such as Zustand, and Jotai

I have not deeply looked into these libraries, but Jotai uses atoms similar to how signals work in preact, and Solid.js, and Zustand is a slimmer version of redux, with more simple syntax.

They offer lighter syntax, and offer more of a library approach rather than a framework.

### Post Script

Lastly, this is a great resource:
Theo - t3․gg - You still use Redux? https://www.youtube.com/watch?v=5-1LM2NySR0

A lot of the state manage on the frontend has changed dynamically in the last few years, but I think it is converging around signals.
