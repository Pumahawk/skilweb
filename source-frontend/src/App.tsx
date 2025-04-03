import { useEffect, useState } from 'react'
import './App.css'
import Dashboard from './dashboard/Dashboard'
import { getSiteMetadata, SiteMetadata } from './services/layout';

function App() {
	const [value, setValue] = useState(undefined as SiteMetadata | undefined);
	useEffect(() => {
		getSiteMetadata().then(metadata => setValue(metadata))
	}, [])
  return (
    <>
    { value ? <Dashboard/> : <div>Loading...</div> }
    </>
  )
}

export default App
