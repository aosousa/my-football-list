import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Title } from '@angular/platform-browser';
import { FlashMessagesService } from 'angular2-flash-messages';

// Services
import { FootballService } from '@services/football.service';

@Component({
    selector: 'reset-password',
    templateUrl: './reset-password.component.html',
    styleUrls: ['./reset-password.component.css']
})
export class ResetPasswordComponent implements OnInit {
    resetPasswordForm: FormGroup;
    submitted = false;

    constructor(
        private _formBuilder: FormBuilder,
        private _titleService: Title,
        private _footballService: FootballService,
        private _flashMessageService: FlashMessagesService,
    ) {
        this.resetPasswordForm = this._formBuilder.group({
            email: ['', [Validators.required, Validators.email, Validators.pattern('^[a-z0-9._%+-]+@[a-z0-9.-]+\\.[a-z]{2,4}$')]]
        });
    }

    ngOnInit() {
        this._titleService.setTitle("Football Tracker - Reset Password")
    }

    // convenience getter for easy access to form fields
    get f() {
        return this.resetPasswordForm.controls;
    }

    sendResetPasswordEmail() {
        this.submitted = true;

        // stop here if the form is invalid
        if (this.resetPasswordForm.invalid) {
            return
        }

        this._footballService.checkEmailExistence(this.resetPasswordForm.value).then(response => {
            // account with email exists - send email
            if (response.rows === 1) {
                this._footballService.sendResetPasswordEmail(this.resetPasswordForm.value).then(response => {
                    if (response.success) {
                        this._flashMessageService.show('Request was successful. An e-mail was sent to the address provided with further instructions to reset your password.', {
                            cssClass: 'alert-success',
                            timeout: 1000000
                        });
                    }
                }).catch(error => {
                    // account doesn't exist - show error message
                    this.submitted = false;
                    this._flashMessageService.show('A reset password request for that e-mail has already been made recently. Check the e-mail address provided for further instructions on how to reset your password.', {
                        cssClass: 'alert-danger',
                        timeout: 1000000
                    });
                });
            } else {
                // account doesn't exist - show error message
                this.submitted = false;
                this._flashMessageService.show('There is no account registered under that e-mail address in the platform.', {
                    cssClass: 'alert-danger',
                    timeout: 1000000
                });
            }
        });
    }
}