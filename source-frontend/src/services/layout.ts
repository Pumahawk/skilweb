
export interface SiteMetadata {
	pages: BackendPage[];
}

export type PageType = 'search';
export type FilterType = 'text';

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

export async function getSiteMetadata(): Promise<SiteMetadata> {
	return new Promise(resolve => {
		const projectSearchPage: SearchPage = {
			type: 'search',
			filters: []
		};
		setTimeout(() => {
			resolve({
				pages: [ projectSearchPage ]
			})
		}, 10000);
	})
}
