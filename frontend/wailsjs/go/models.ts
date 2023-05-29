export namespace main {
	
	export class Mod {
	    name: string;
	    version: string;
	    fileName: string;
	
	    static createFrom(source: any = {}) {
	        return new Mod(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.version = source["version"];
	        this.fileName = source["fileName"];
	    }
	}

}

