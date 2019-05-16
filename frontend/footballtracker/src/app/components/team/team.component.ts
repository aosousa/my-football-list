import { Component, OnInit, AfterViewInit, OnDestroy, ViewChild } from '@angular/core';
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
    selector: 'team',
    templateUrl: './team.component.html',
    styleUrls: ['./team.component.css']
})
export class TeamComponent implements OnInit, AfterViewInit, OnDestroy {
    @ViewChild(DataTableDirective)
    dtElement: DataTableDirective;
  
    user: any = {};
    public teamFixtures: any = [];
    sessionUserId: number;
    teamId: number;
    dtOptions: DataTables.Settings = {};
    dtTrigger: Subject<any> = new Subject();

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessagesService: FlashMessagesService,
        private _route: ActivatedRoute,
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
        this.teamId = Number(this._route.snapshot.paramMap.get('id'));
        this.loadTeamFixtures(this.teamId);
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
     * Load a team's fixtures stored in the database
     * @param {number} teamID ID of the team
     */
    loadTeamFixtures(teamID: number) {
        this._footballService.getTeamFixtures(teamID).then(response => {
            if (response.success && response.rows > 0) {
                this.teamFixtures = response.data;
                this._titleService.setTitle("Football Tracker - " + this.teamFixtures.team.name);

                this.dtElement.dtInstance.then((dtInstance: DataTables.Api) => {
                    // Destroy the table first
                    dtInstance.destroy();

                    // Call the dtTrigger to re-render again
                    this.dtTrigger.next();
                });
            } else {
                this.teamFixtures = response.data;

                // team exists but has no fixtures
                this._flashMessagesService.show("No matches found", {
                    cssClass: 'alert-danger',
                    timeout: 10000000000000
                });
            }
        }).catch((error: Response) => {
            this._flashMessagesService.show(error.json().error, {
                cssClass: 'alert-danger',
                timeout: 10000000000000
            });
        });
    }

    /**
     * Set fixture status as "watched" or "want to watch"
     * @param {number} fixtureID ID of the fixture in the platform
     * @param {number} fixtureStatus Fixture status (1 = Watched, 2 = Want to Watch)
     * @param {number} userFixtureId ID of the user fixture row (0 means a new one will be created, otherwise it's an update to the row with the ID received)
     * @param {number} position Position of the fixture in the teamFixtures array
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
                this.teamFixtures.fixtures[position].userFixtureStatus = fixtureStatus
                this.teamFixtures.fixtures[position].userFixtureID = response.data.userFixtureId
            }
        })
    }

    /**
     * Delete a user fixture status row
     * @param {number} userFixtureId ID of the user fixture row to delete
     * @param {number} position Position of the fixture in the teamFixtures array
     */
    deleteUserFixture(userFixtureId: number, position: number) {
        this._footballService.deleteUserFixture(userFixtureId).then(response => {
            if (response.success) {
                this.teamFixtures.fixtures[position].userFixtureStatus = 0
                this.teamFixtures.fixtures[position].userFixtureID = 0
            }
        })
    }
}