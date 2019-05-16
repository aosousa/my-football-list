import { Component, OnInit, ViewChild, AfterViewInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';
import { Response } from '@angular/http';
import { Subject } from 'rxjs';
import { DataTableDirective } from 'angular-datatables';

// 3rd party
import * as _ from 'lodash';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'league',
    templateUrl: './league.component.html',
    styleUrls: ['./league.component.css']
})
export class LeagueComponent implements OnInit, AfterViewInit, OnDestroy {
    @ViewChild(DataTableDirective)
    dtElement: DataTableDirective;

    user: any = {}
    public leagueFixtures: any = [];
    sessionUserId: number;
    leagueId: number;
    dtOptions: DataTables.Settings = {};
    dtTrigger: Subject<any> = new Subject();

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessagesService: FlashMessagesService,
        private _route: ActivatedRoute
    ) { }

    ngOnInit() {
        this.dtOptions = {
            ordering: false,
            paging: true,
            pagingType: 'full_numbers',
            pageLength: 10,
            processing: true,
            lengthChange: false,
        };

        this.sessionUserId = Number(sessionStorage.getItem('userId'));
        this.leagueId = Number(this._route.snapshot.paramMap.get('id'));
        this.loadLeagueFixtures(this.leagueId);
        this._footballService.getUser(this.sessionUserId).then(response => {
            if (response.success) {
                this.user = response.data;
            }
        });
    }

    ngAfterViewInit(): void {
        this.dtTrigger.next();
    }

    ngOnDestroy(): void {
        // Do not forget to unsubscribe the event
        this.dtTrigger.unsubscribe();
    }

    /**
     * Load a league's fixtures stored in the database
     * @param {number} leagueID ID of the league
     */
    loadLeagueFixtures(leagueID: number) {
        this._footballService.getLeagueFixtures(leagueID).then(response => {
            if (response.success && response.rows > 0) {
                this.leagueFixtures = response.data;
                this._titleService.setTitle("Football Tracker - " + this.leagueFixtures.league.name);

                this.dtElement.dtInstance.then((dtInstance: DataTables.Api) => {
                    // Destroy the table first
                    dtInstance.destroy();

                    // Call the dtTrigger to re-render again
                    this.dtTrigger.next();
                });
            } else {
                this.leagueFixtures = response.data;

                // league exists but has no fixtures
                this._flashMessagesService.show("No matches found", {
                    cssClass: 'alert-danger',
                    timeout: 10000000000000
                })
            }
        }).catch((error: Response) => {
            this._flashMessagesService.show(error.json().error, {
                cssClass: 'alert-danger',
                timeout: 10000000000000
            });
        })
    }

    /**
     * Set fixture status as "watched" or "want to watch"
     * @param {number} fixtureID ID of the fixture in the platform
     * @param {number} fixtureStatus Fixture status (1 = Watched, 2 = Want to Watch)
     * @param {number} userFixtureId ID of the user fixture row (0 means a new one will be created, otherwise it's an update to the row with the ID received)
     * @param {number} position Position of the league in the leagueFixtures array
     */
    setUserFixtureStatus(fixtureID: number, fixtureStatus: number, userFixtureId: number, position: number) {
        let userFixtureID = userFixtureId == 0 ? null : userFixtureId

        let userFixtureStatus = {
            fixtureId: fixtureID,
            status: fixtureStatus,
            userFixtureId: userFixtureID
        }

        this._footballService.createUserFixture(userFixtureStatus).then(response => {
            if (response.success) {
                this.leagueFixtures.fixtures[position].userFixtureStatus = fixtureStatus
                this.leagueFixtures.fixtures[position].userFixtureID = response.data.userFixtureId            
            }
        })
    }

    /**
     * Delete a user fixture status row
     * @param {number} userFixtureId ID of the user fixture row to delete
     * @param {number} position Position of the fixture in the leagueFixtures array
     */
    deleteUserFixture(userFixtureId: number, position: number) {
        this._footballService.deleteUserFixture(userFixtureId).then(response => {
            if (response.success) {
                this.leagueFixtures.fixtures[position].userFixtureStatus = 0
                this.leagueFixtures.fixtures[position].userFixtureID = 0       
            }
        })
    }
}