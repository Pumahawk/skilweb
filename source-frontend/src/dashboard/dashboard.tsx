import { BackendPage, SiteMetadata } from "../services/metadata/dto"

export function Dashboard(props: {metadata: SiteMetadata}) {
	return (
		<div>
			<Header/>
			<div>
				<Menu pages={props.metadata.pages}/>
				<Page/>
			</div>
			<Footer/>
		</div>
	)
}

export function Header() {
	return (
		<div>Header</div>
	)
}

export function Menu(props: {pages: BackendPage[]}) {
	return (
		<div>
			<div>Menu</div>
			{ props.pages && props.pages.map((p, i) => (<div id={i.toString()}>Type: {p.type}</div>)) }
		</div>
	)
}

export function Page() {
	return (
		<div>Page</div>
	)
}

export function Footer() {
	return (
		<div>Footer</div>
	)
}
