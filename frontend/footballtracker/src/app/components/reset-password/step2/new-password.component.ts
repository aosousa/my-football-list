import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// Services 
import { FootballService } from '@services/football.service';

// import custom validator to validate that password and confirm password fields match
import { MustMatch } from '@helpers/must-match.validator';

@Component({
    selector: 'new-password',
    templateUrl: './new-password.component.html',
    styleUrls: ['./new-password.component.css']
})
export class NewPasswordComponent implements OnInit {
    newPasswordForm: FormGroup;
    submitted = false;
    invalidToken = true;
    changedSuccessfully = false;
    token: string;

    constructor(
        private _formBuilder: FormBuilder,
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
        private _route: ActivatedRoute,
        private _router: Router
    ) {
        this.newPasswordForm = this._formBuilder.group({
            password: ['', [Validators.required, Validators.minLength(6)]],
            confirmPassword: ['', Validators.required],
        }, {
            validator: MustMatch('password', 'confirmPassword')
        });
    }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Reset Password");
        
        // check if token is still valid
        this.token = this._route.snapshot.paramMap.get('token');
        this._footballService.validateResetPasswordToken(this.token).then(response => {
            this.invalidToken = false;
        }).catch(error => {
            this.invalidToken = true;
            this._flashMessageService.show('Reset password token is no longer valid. You must request a new password reset in order to change your password.', {
                cssClass: 'alert-danger',
                timeout: 1000000
            });
        });
    }

    // convenience getter for easy access to form fields
    get f() {
        return this.newPasswordForm.controls;
    }

    resetPassword() {
        this.submitted = true;

        // stop here if form is invalid
        if (this.newPasswordForm.invalid) {
            return
        }

        if (!this.changedSuccessfully) {
            let passwordResetInfo = this.newPasswordForm.value;
            passwordResetInfo.token = this.token
    
            this._footballService.resetPassword(passwordResetInfo).then(response => {
                if (response.success) {
                    this.submitted = false;
                    this.invalidToken = false;
                    this.changedSuccessfully = true;
                    this._flashMessageService.show('Password was changed successfully. Redirecting you to login page in 5 seconds.', {
                        cssClass: 'alert-success',
                        timeout: 5000
                    });
                    setTimeout(() => {
                        this._router.navigate(['/login'])
                    }, 5000);
                }
            }).catch(error => {
                this.invalidToken = true;
                this._flashMessageService.show('Reset password token is no longer valid. You must request a new password reset in order to change your password.', {
                    cssClass: 'alert-danger',
                    timeout: 1000000
                });        
            });
        }
    }
}