import { SiteMetadata } from "../services/metadata/dto"

export function Dashboard(props: {metadata: SiteMetadata}) {
	return (
		<div>
			<Header/>
			<div>
				<Menu/>
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

export function Menu() {
	return (
		<div>Menu</div>
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
