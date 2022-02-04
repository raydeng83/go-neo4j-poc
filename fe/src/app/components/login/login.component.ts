import {Component, OnInit} from '@angular/core';
import {AppConst} from "../../constants";
import {HttpClient, HttpHeaders} from "@angular/common/http";
import {Observable} from "rxjs";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.scss']
})
export class LoginComponent implements OnInit {
  backend = AppConst.BACKEND_URL;
  email;
  password;

  constructor(
    private http: HttpClient
  ) {
  }

  login() {
    const credentials = {
      'user': {
        'email': this.email,
        'password': this.password
      }
    }
    this.sendCredentials(credentials).subscribe(
      res => {
        console.log(res)
        // localStorage.setItem(AppConst.TOKEN_NAME, )
      }, error => {

      }
    )
  }

  sendCredentials(credentials): Observable<any> {
    const url = this.backend + '/users/login';
    const httpOptions = {
      headers: new HttpHeaders({'Content-Type': 'application/json'})
    };
    return this.http.post(url, credentials, httpOptions);
  }

  ngOnInit(): void {
  }

}
