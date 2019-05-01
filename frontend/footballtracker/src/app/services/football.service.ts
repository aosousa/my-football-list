import { Injectable } from '@angular/core';
import { Http, Headers, Response } from '@angular/http';

// RxJS imports
import { BehaviorSubject } from 'rxjs';
import { map, timeout } from 'rxjs/operators';

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
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    // Fixture methods

    /**
     * Get fixtures in a given date
     * @param {string} date Date in YYYY-mm-dd format
     */
    getFixturesByDate(date: string) {
        return this._http.get(`${this.config_link}/fixtures/by-date/${date}`, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Get date/time of the last fixture update in YYYY-mm-dd format
     */
    getLastFixtureUpdate() {
        return this._http.get(`${this.config_link}/fixtures/last-update`, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    // User methods

    /**
     * Get a user's information
     * @param {number} id Id of the user
     */
    getUser(id: number) {
        return this._http.get(`${this.config_link}/users/${id}`, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Update a user's information
     * @param {number} id Id of the user
     * @param {any} userinfo New information for the user
     */
    updateUser(id: number, userinfo: any) {
        return this._http.put(`${this.config_link}/users/${id}`, userinfo, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    // User/Fixture methods

    /**
     * Get a user's fixtures
     * @param {number} id Id of the user
     */
    getUserFixtures(id: number) {
        return this._http.get(`${this.config_link}/users/${id}/fixtures`, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Create a user_fixture row
     * @param {any} fixtureInfo Information about the fixture (id, status)
     */
    createUserFixture(fixtureInfo: any) {
        return this._http.post(`${this.config_link}/users/fixtures`, fixtureInfo, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Delete a user_fixture row
     * @param {number} id Id of the user fixture row
     */
    deleteUserFixture(id: number) {
        return this._http.delete(`${this.config_link}/user-fixtures/${id}`, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    // Authentication methods

    /**
     * Log the user in
     * @param {any} userInfo Information required to perform login
     * @returns {Promise<any>}
     */
    login(userInfo: any) {
        return this._http.post(`${this.config_link}/login`, userInfo, this.options)
            .pipe(
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
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Check if a username is taken
     * @param {any} userInfo Object with username
     * @returns {Promise<any>}
     */
    checkUsernameExistence(userInfo: any) {
        return this._http.post(`${this.config_link}/users/username-existence`, userInfo, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Check if an email is taken
     * @param {any} userInfo Object with email
     * @returns {Promise<any>}
     */
    checkEmailExistence(userInfo: any) {
        return this._http.post(`${this.config_link}/users/email-existence`, userInfo, this.options)
            .pipe(
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

    /**
     * Get information of the logged in user
     */
    getCurrentUser(): Promise<any> {
        return this._http.get(`${this.config_link}/users/current`, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /** 
     * Checks if a reset password token is still valid
     * @param {string} token Reset password token
     * @returns {Promise<any>}
     */
    validateResetPasswordToken(token: string) {
        return this._http.get(`${this.config_link}/tokens/${token}/valid`, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Reset a user's password
     * @param {any} passwordInfo Reset password information
     * @returns {Promise<any>}
     */
    resetPassword(passwordInfo: any) {
        return this._http.post(`${this.config_link}/reset-password`, passwordInfo, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Change a user's password
     * @param {number} userId Id of the user that is changing password
     * @param {any} passwordInfo Information required to perform a password change
     */
    changePassword(userId: number, passwordInfo: any) {
        return this._http.put(`${this.config_link}/users/${userId}/change-password`, passwordInfo, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    // Contact methods

    /**
     * Send an email through contact form
     * @param {any} emailInfo Object with required information (type, subject, message)
     * @returns {Promise<any>}
     */
    sendEmail(emailInfo: any) {
        return this._http.post(`${this.config_link}/contact`, emailInfo, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }

    /**
     * Send a reset password email 
     * @param {any} emailInfo Email information
     * @returns {Promise<any>}
     */
    sendResetPasswordEmail(emailInfo: any) {
        return this._http.post(`${this.config_link}/reset-password-email`, emailInfo, this.options)
            .pipe(
                timeout(5000),
                map((resp: Response) => resp.json())
            ).toPromise();
    }
}