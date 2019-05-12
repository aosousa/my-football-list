import { Component, OnInit, AfterViewInit, OnDestroy, ViewChildren, QueryList } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// Services
import { FootballService } from '@services/football.service';
import { DataTableDirective } from 'angular-datatables';
import { Subject } from 'rxjs';

@Component({
    selector: 'profile',
    templateUrl: './profile.component.html',
    styleUrls: ['./profile.component.css']
})
export class ProfileComponent implements OnInit, AfterViewInit, OnDestroy {
    @ViewChildren(DataTableDirective)
    dtElements: QueryList<DataTableDirective>;

    user: any = {};
    userId: number;
    sessionUserId: number;
    userExists = true;
    dtOptions: DataTables.Settings = {};
    dtTrigger: Subject<any> = new Subject();
    public fixturesWatched: any = [];
    public fixturesInterestedIn: any = [];

    constructor(
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
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
            destroy: true
        };

        this._titleService.setTitle("Football Tracker");
        this.userId = Number(this._route.snapshot.paramMap.get('id'));
        this.sessionUserId = Number(sessionStorage.getItem('userId'));
        this.loadUserFixtures(this.userId);

        this._footballService.getUser(this.userId).then(response => {
            if (response.success) {
                this.user = response.data;
                this._titleService.setTitle("Football Tracker - " + this.user.username);
            } 
        }).catch(error => {
            this.userExists = false;
            this._flashMessageService.show('User does not exist.', {
                cssClass: 'alert-danger',
                timeout: 1000000
            });
        });
    }

    ngAfterViewInit(): void {
        this.dtTrigger.next();
    }

    ngOnDestroy(): void {
        // Do not forget to unsubscribe the event
        this.dtTrigger.unsubscribe();
    }

    rerender(): void {
        this.dtElements.forEach((dtElement: DataTableDirective) => {
            dtElement.dtInstance.then((dtInstance: DataTables.Api) => {
                // Destroy the table first
                dtInstance.destroy();

                // Call the dtTrigger to re-render again
                this.dtTrigger.next();
            });
        });
    }

    loadUserFixtures(userID: number) {
        this._footballService.getUserFixtures(userID).then(response => {
            if (response.success) {
                this.fixturesWatched = response.data.watched;
                this.fixturesInterestedIn = response.data.interestedIn;
                this.rerender();
            }
        }).catch(error => {
            this._flashMessageService.show("Error getting user's fixtures. Please try again later.", {
                cssClass: 'alert-danger',
                timeout: 10000000000000
            });
        })
    }
}