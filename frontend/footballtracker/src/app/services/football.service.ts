import { Injectable } from '@angular/core';
import { Http, Headers, Response } from '@angular/http';

// RxJS imports
import { BehaviorSubject } from 'rxjs';
import { map, timeout, retry } from 'rxjs/operators';

// Models
import { Config } from '@models/config';

// Services
import { ConfigService } from '@services/config.service';

@Injectable()
export class FootballService {
    static config: Config;

    private config_link;
    private options;
    private messageSource = new BehaviorSubject('');
    private usernameSource = new BehaviorSubject('');

    currentMessage = this.messageSource.asObservable();
    usernameMessage = this.usernameSource.asObservable();

    constructor(
        private _http: Http, 
        private _configService: ConfigService
    ) {
        // get configuration from JSON file
        FootballService.config = ConfigService.getConfiguration();
    }
}