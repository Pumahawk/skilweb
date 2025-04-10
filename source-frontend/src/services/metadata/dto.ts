export type PageType = 'search';
export type FilterType = 'text';

export interface SiteMetadata {
	pages: BackendPage[];
}

export interface BackendPage {
	type: PageType;
}

export interface SearchFilterData {
	name: string;
	type: FilterType;
}

export interface TextFilter extends SearchFilterData {
	name: string;
	type: 'text';
	label: string;
}

export interface SearchPageData extends BackendPage {
	type: 'search';
	filters: SearchFilterData[];
}
