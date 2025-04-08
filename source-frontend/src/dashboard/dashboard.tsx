import { Drawer, List, ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material"
import { BackendPage, SiteMetadata } from "../services/metadata/dto"
import { People, Work } from "@mui/icons-material"

import Box from "@mui/material/Box"

interface VoiceMenu {
	id: string;
	icon: MatIcon;
	text: string;
}

export function Dashboard(props: {metadata: SiteMetadata}) {
	const voicesMenu: VoiceMenu[] = [
		{ id: "01", icon: "people", text: "Users" },
		{ id: "02", icon: "work", text: "Projects" },
	]
	return (
		<Box display="flex" position="fixed">
			<Drawer variant="permanent">
				<List sx={{width: 250}}>
					{ voicesMenu.map(m => (
						<ListItem key={m.id} disablePadding>
							<ListItemButton>
								<ListItemIcon>
									<Icon icon={m.icon}/>
								</ListItemIcon>
								<ListItemText primary={m.text}/>
							</ListItemButton>
						</ListItem>
					))}
				</List>
			</Drawer>
			<Box flex="1">
			main
			</Box>
		</Box>
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

type MatIcon = "people" | "work"

function Icon(props: {icon: MatIcon}) {
	switch (props.icon) {
		case "people":
			return <People/>
		case "work":
			return <Work/>
	}
}
