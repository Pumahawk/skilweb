import { useEffect, useState } from 'react'
import './App.css'
import { metadataClient } from './services/metadata/client';
import { SiteMetadata } from './services/metadata/dto';
import { Dashboard } from './dashboard/dashboard';

function App() {
	const [value, setValue] = useState(undefined as SiteMetadata | undefined);
	const [error, setError] = useState(undefined as any | undefined);
	useEffect(() => {
		metadataClient.getSiteMetadata()
		.then(metadata => setValue(metadata))
		.catch(error => {
			console.log("Unable to retrieve metadata informations", error)
			setError(error)
			
		});
	}, [])
  return (
    <>
    { error && <MetadataError error={error}/>}
    { error == undefined && (value ? <Dashboard metadata={value}/> : <div>Loading...</div>) }
    </>
  )
}

function MetadataError(props: {error: any}) {
	return (
		<>
		Error...
		</>
	)
}

export default App
