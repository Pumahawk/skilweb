import { Drawer, List, ListItem, ListItemButton, ListItemIcon, ListItemText } from "@mui/material"
import { BackendPage, SearchFilterData, SearchPageData, SiteMetadata } from "../services/metadata/dto"
import { People, Work } from "@mui/icons-material"

import Box from "@mui/material/Box"

const drawerWidth = 240;

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
	const searchData: SearchPageData = {
		type: "search",
		filters: [{type: "text", name: "Testing..."}],
	}
	return (
		<Box sx={{display: "flex"}}>
			<Menu voicesMenu={voicesMenu}/>
			<SearchPage data={searchData}/>
		</Box>
	)
}

export function Header() {
	return (
		<div>Header</div>
	)
}

export function Menu(props: {voicesMenu: VoiceMenu[]}) {
	return (
		<Drawer
		sx={{
		  width: drawerWidth,
		  flexDirection: "column",
		  flexShrink: 0,
		  '& .MuiDrawer-paper': {
		    width: drawerWidth,
		    boxSizing: 'border-box',
		  },
		}}
		variant="permanent">
			<List sx={{width: "100%"}}>
				{ props.voicesMenu.map(m => (
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

function SearchPage(props: {data: SearchPageData}) {
	return (
	    	<Box
			component="main"
			sx={{ flexGrow: 1, bgcolor: 'background.default', p: 3, flexDirection: "column"}}
	      	>
			<SearchFilters filters={props.data.filters}/>
			<SearchContent/>
		</Box>
	)
}

function SearchFilters(props: {filters: SearchFilterData[]}) {
	return (
		<Box>
			Filters....
		</Box>
	)
}

function SearchContent() {
	return (
		<Box>
			Content...
		</Box>
	)
}
