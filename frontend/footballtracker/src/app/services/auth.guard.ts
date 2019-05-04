import { ActivatedRouteSnapshot, CanActivate, Router, RouterStateSnapshot } from '@angular/router';
import { Injectable }  from '@angular/core';

// Services
import { FootballService } from '@services/football.service';

@Injectable()
export class AuthGuard implements CanActivate {
    constructor(
        private _footballService: FootballService,
        private _router: Router 
    ) { }

    canActivate(route: ActivatedRouteSnapshot, state: RouterStateSnapshot): boolean {
        const loginStatus = this._footballService.isAuthenticated();
        if (loginStatus >= 0) {
            this._footballService.changeMessage('true');
            return true;
        } else {
            this._footballService.changeMessage('false');
            this._router.navigate(['/']);
            return false;
        }
    }
}