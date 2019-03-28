import { map } from 'rxjs/operators';
import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

// Models
import { Config } from '@models/config';

@Injectable()
export class ConfigService {
    static config: Config;

    constructor(private _http: Http) { }

    /**
     * Get application configuration
     * @returns {Config}
     */
    static getConfiguration(): Config {
        return ConfigService.config;
    }

    /**
     * Load application configuration 
     * @param {string} url Path / URL of the config file
     */
    loadConfig(url: string) {
        return new Promise((resolve) => {
            this._http.get(url).pipe(map(res => res.json()))
                .subscribe(config => {
                    ConfigService.config = config;
                    resolve();
                });
        });
    }
}