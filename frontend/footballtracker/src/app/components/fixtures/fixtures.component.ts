import { Component, OnInit } from '@angular/core';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';
import { NgbDateStruct } from '@ng-bootstrap/ng-bootstrap';

// 3rd party
import * as _ from 'lodash';

// Services
import { FootballService } from '@services/football.service';
import { UtilsService } from '@services/utils.service';

@Component({
    selector: 'fixtures',
    templateUrl: './fixtures.component.html',
    styleUrls: ['./fixtures.component.css']
})
export class FixturesComponent implements OnInit {
    user: any = {};
    fixtures: any;
    groupedFixtures: any;
    dateVar: string;
    lastFixtureUpdate: string
    sessionUserId: number;
    dateModel: NgbDateStruct;
    date: {
        year: number,
        month: number,
        day: number
    };

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _utilsService: UtilsService,
        private _flashMessageService: FlashMessagesService,
    ) { }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Fixtures");
        this.sessionUserId = Number(sessionStorage.getItem('userId'));

        const now = new Date();
        this.dateVar = this._utilsService.buildDate(now.getFullYear(), _.padStart(String(now.getMonth() + 1), 2, '0'), _.padStart(String(now.getDate()), 2, '0'));

        this.loadFixtures(this.dateVar);
        this.getLastFixtureUpdate();

        this._footballService.getUser(this.sessionUserId).then(response => {
            if (response.success) {
                this.user = response.data;
            }
        });
    }

    /**
     * Loads fixtures from a given day
     * @param {string} date Date in YYYY-mm-dd format
     */
    loadFixtures(date: string) {
        this._footballService.getFixturesByDate(date).then(response => {
            if (response.success) {
                this.fixtures = response.data;
            }

            // group fixtures by league
            this.groupedFixtures = _(this.fixtures)
                .groupBy(x => x.league.leagueId)
                .map((fixtures, league) => ({fixtures, league}))
                .value();
        }).catch(error => {
            this._flashMessageService.show('An error occurred while updating the fixtures list.', {
                cssClass: 'alert-danger',
                timeout:  5000
            });
        });
    }

    /**
     * Get date/time of last fixture update
     */
    getLastFixtureUpdate() {
        this._footballService.getLastFixtureUpdate().then(response => {
            if (response.success) {
                this.lastFixtureUpdate = response.data;
            }
        })
    }
    
    /**
     * Set fixture status as "watched" or "want to watch"
     * @param {number} fixtureID ID of the fixture in the platform
     * @param {number} fixtureStatus Fixture status (1 = Watched, 2 = Want to Watch)
     * @param {number} userFixtureId ID of the user fixture row (0 means a new one will be created, otherwise it's an update to the row with the ID received)
     * @param {number} leaguePosition Position of the league in the groupedFixtures array
     * @param {number} fixturePosition Position of the fixture in the groupedFixtures array
     */
    setUserFixtureStatus(fixtureID: number, fixtureStatus: number, userFixtureId: number, leaguePosition: number, fixturePosition: number) {
        let userFixtureID = userFixtureId == 0 ? null : userFixtureId

        let userFixtureStatus = {
            fixtureId: fixtureID,
            status: fixtureStatus,
            userFixtureId: userFixtureID
        }

        this._footballService.createUserFixture(userFixtureStatus).then(response => {
            if (response.success) {
                this.groupedFixtures[leaguePosition].fixtures[fixturePosition].userFixtureID = response.data.userFixtureId
                this.groupedFixtures[leaguePosition].fixtures[fixturePosition].userFixtureStatus = fixtureStatus
            }
        })
    }

    /**
     * Delete a user fixture status row
     * @param {number} userFixtureId ID of the user fixture row to delete
     * @param {number} leaguePosition Position of the league in the groupedFixtures array
     * @param {number} fixturePosition Position of the fixture in the groupedFixtures array
     */
    deleteUserFixture(userFixtureId: number, leaguePosition: number, fixturePosition: number) {
        this._footballService.deleteUserFixture(userFixtureId).then(response => {
            if (response.success) {
                this.groupedFixtures[leaguePosition].fixtures[fixturePosition].userFixtureID = 0
                this.groupedFixtures[leaguePosition].fixtures[fixturePosition].userFixtureStatus = 0
            }
        })
    }

    /**
     * Update groupedFixtures array to one with the fixtures from the chosen date
     */
    filterFixturesByDate() {
        let filterDate = this._utilsService.buildDate(this.dateModel.year, _.padStart(String(this.dateModel.month), 2, '0'), _.padStart(String(this.dateModel.day), 2, '0'));
        
        this._footballService.getFixturesByDate(filterDate).then(response => {
            if (response.success) {
                if (response.success) {
                    this.fixtures = response.data;
                }
    
                // group fixtures by league
                this.groupedFixtures = _(this.fixtures)
                    .groupBy(x => x.league.leagueId)
                    .map((fixtures, league) => ({fixtures, league}))
                    .value();
            }
        }).catch(error => {

        });
    }
}