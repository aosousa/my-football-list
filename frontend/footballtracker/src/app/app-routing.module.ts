import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

// Components
import { FixturesComponent } from '@components/fixtures/fixtures.component';
import { IndexComponent } from '@components/index/index.component';
import { LoginComponent } from '@components/login/login.component';
import { SignupComponent } from '@components/signup/signup.component';

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
        path: 'fixtures',
        component: FixturesComponent,
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