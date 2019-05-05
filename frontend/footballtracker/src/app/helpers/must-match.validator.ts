import { FormGroup, ValidatorFn, AbstractControl, ValidationErrors } from '@angular/forms';

// custom validator to check that two fields match
export function MustMatch(controlName: string, matchingControlName: string) {
    return (formGroup: FormGroup) => {
        const control = formGroup.controls[controlName];
        const matchingControl = formGroup.controls[matchingControlName];

        if (matchingControl.errors && !matchingControl.errors.mustMatch) {
            // return if another validator has already found an error on the matchingControl
            return;
        }

        // set error on matchingControl if validation fails
        if (control.value !== matchingControl.value) {
            matchingControl.setErrors({ mustMatch: true });
        } else {
            matchingControl.setErrors(null);
        }
    }
}

export const matchingInputsValidator: ValidatorFn = (control: AbstractControl) : ValidationErrors | null => {
    let pass;
    control.root.get('password') ? pass = control.root.get('password').value : pass = null;
    let confirmPass = control.value;
    return pass === confirmPass ? null : { mustMatch: false };
}