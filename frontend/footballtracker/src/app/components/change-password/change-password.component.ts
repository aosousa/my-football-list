import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';
import { Response } from '@angular/http';

// Services
import { FootballService } from '@services/football.service';

// import custom validator to validate that password and confirm password fields match
import { MustMatch } from '@helpers/must-match.validator';

@Component({
    selector: 'change-password',
    templateUrl: './change-password.component.html',
    styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent implements OnInit {
    changePasswordForm: FormGroup;
    userId: number;
    sessionUserId: number;
    submitted = false;
    changedSuccessfully = false;
    canEdit = true;
    processing = false;

    constructor(
        private _formBuilder: FormBuilder,
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
        private _route: ActivatedRoute,
        private _router: Router
    ) {
        this.changePasswordForm = this._formBuilder.group({
            currentPassword: ['', Validators.required],
            newPassword: ['', [Validators.required, Validators.minLength(6)]],
            confirmNewPassword: ['', Validators.required]
        }, {
            validator: MustMatch('newPassword', 'confirmNewPassword')
        })
    }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Change Password")
        this.userId = Number(this._route.snapshot.paramMap.get('id'));
        this.sessionUserId = Number(sessionStorage.getItem('userId'));

        if (this.userId != this.sessionUserId) {
            this.canEdit = false;
            this._flashMessageService.show('You are not authorized to perform this action.', {
                cssClass: 'alert-danger',
                timeout: 100000000
            });
        }
    }

    // convenience getter for easy access
    get f() {
        return this.changePasswordForm.controls;
    }

    /**
     * Change user's password.
     * Show validation errors if form submission is invalid
     */
    changePassword() {
        this.submitted = true;
        this.processing = true;

        // stop here if form is invalid
        if (this.changePasswordForm.invalid) {
            this.processing = false;
            return
        }

        if (!this.changedSuccessfully) {
            this._footballService.changePassword(this.userId, this.changePasswordForm.value).then(response => {
                this.submitted = false;
                this.processing = false;
                if (response.success) {
                    this._flashMessageService.show('Your password was changed successfully. Redirecting you to your profile page in 5 seconds.', {
                        cssClass: 'alert-success',
                        timeout: 5000
                    });
                    setTimeout(() => {
                        this._router.navigate(['/user/' + this.userId])
                    }, 5000);
                }
            }).catch((error: Response) => {
                this.processing = false;
                this._flashMessageService.show(error.json().error, {
                    cssClass: 'alert-danger',
                    timeout: 10000
                });
            });
        }
    }
}