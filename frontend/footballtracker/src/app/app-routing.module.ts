import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

// Components
import { AboutComponent } from '@components/about/about.component';
import { ContactComponent } from '@components/contact/contact.component';
import { FixturesComponent } from '@components/fixtures/fixtures.component';
import { IndexComponent } from '@components/index/index.component';
import { LoginComponent } from '@components/login/login.component';
import { PrivacyPolicyComponent } from '@components/privacy-policy/privacy-policy.component';
import { ProfileComponent } from '@components/profile/profile.component';
import { SignupComponent } from '@components/signup/signup.component';
import { TermsOfServiceComponent } from '@components/terms-of-service/terms-of-service.component';

// Services
import { AuthGuard } from '@services/auth.guard';

const routes: Routes = [
    {
        path: '',
        component: IndexComponent,
    },
    {
        path: 'signup',
        component: SignupComponent,
    },
    {
        path: 'login',
        component: LoginComponent
    },
    {
        path: 'about',
        component: AboutComponent
    },
    {
        path: 'privacy-policy',
        component: PrivacyPolicyComponent
    },
    {
        path: 'terms-of-service',
        component: TermsOfServiceComponent
    },
    {
        path: 'fixtures',
        component: FixturesComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'contact',
        component: ContactComponent,
        canActivate: [AuthGuard]
    },
    {
        path: 'user/:id',
        component: ProfileComponent,
        canActivate: [AuthGuard]
    }
]

@NgModule({
    imports: [
        RouterModule.forRoot(routes, {
            useHash: false,
            anchorScrolling: 'enabled'
        })
    ],
    exports: [
        RouterModule
    ]
})
export class AppRoutingModule{}