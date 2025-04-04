export type PageType = 'search';
export type FilterType = 'text';

export interface SiteMetadata {
	pages: BackendPage[];
}

export interface SiteMetadata {
	pages: BackendPage[];
}

export interface BackendPage {
	type: PageType;
}

export interface SearchFilter {
	name: string;
	type: FilterType;
}

export interface TextFilter extends SearchFilter {
	name: string;
	type: 'text';
	label: string;
}

export interface SearchPage extends BackendPage {
	type: 'search';
	filters: SearchFilter[];
}
