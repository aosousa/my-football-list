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

        // get production or development environment settings depending on the app configuration
        this.config_link = (FootballService.config.prod) ? FootballService.config.prodDomain : FootballService.config.devDomain;

        // set request headers
        const headers: Headers = new Headers();
        headers.append("Content-Type", "application/json");
        headers.append("Cache-Control", "no-cache");
        headers.append("Pragma", "no-cache");

        this.options = {
            headers: headers,
            withCredentials: true
        };
    }

    changeMessage(message: string) {
        this.messageSource.next(message);
    }

    changeUsernameSource(username: string) {
        this.usernameSource.next(username);
    }

    // League methods

    /**
     * Get all league records in the database
     * @returns {Promise<any>}
     */
    getAllLeagues() {
        return this._http.get(`${this.config_link}/leagues`, this.options)
            .pipe(
                retry(3),
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /** 
     * Get a league's fixtures
     * @param {number} id ID of the league
     * @returns {Promise<any>}
     */
    getLeagueFixtures(id: number) {
        return this._http.get(`${this.config_link}/leagues/${id}/fixtures`, this.options)
            .pipe(
                retry(3),
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    // Team methods

    /**
     * Get a team's fixtures
     * @param {number} id ID of the team
     * @returns {Promise<any>}
     */
    getTeamFixtures(id: number) {
        return this._http.get(`${this.config_link}/teams/${id}/fixtures`, this.options)
            .pipe(
                retry(3),
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    // User / Authentication methods

    /**
     * Log the user in
     * @param {any} userInfo Information required to perform login
     * @returns {Promise<any>}
     */
    login(userInfo: any) {
        return this._http.post(`${this.config_link}/login`, userInfo, this.options)
            .pipe(
                retry(3),
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Log user out of the platform
     * @returns {Promise<any>}
     */
    logout() {
        return this._http.post(`${this.config_link}/logout`, {}, this.options)
            .pipe(
                retry(3),
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Register user in the platform
     * @param {any} signupInfo Information required to register user in the platform
     * @returns {Promise<any>}
     */
    signup(signupInfo: any) {
        return this._http.post(`${this.config_link}/signup`, signupInfo, this.options)
            .pipe(
                retry(3),
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Check if the user is authenticated in the platform
     * @returns {number}
     */
    isAuthenticated(): number {
        return document.cookie.indexOf("session-token")
    }
}