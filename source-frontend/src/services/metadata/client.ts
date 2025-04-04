import { SiteMetadata } from './dto'

const paths = {
	metadata: "/metadata",
}

export interface ClientConf {
	baseUri: string;
}

export class MetadataClient {
	conf: ClientConf;

	constructor(conf: ClientConf) {
		this.conf = conf;
	}

	async getSiteMetadata(): Promise<SiteMetadata> {
		const resp = await fetch(this.conf.baseUri + paths.metadata)
		if (resp.ok) {
			const body: SiteMetadata = await resp.json()
			return body;
		} else {
			throw new Error(`Error Response: ${resp.statusText}`)
		}
	}
}

export const metadataClient = new MetadataClient({
	baseUri: "/api",
});
