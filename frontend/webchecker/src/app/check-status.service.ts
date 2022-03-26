import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { WebsiteCheckerResponse } from './website-status';

@Injectable({
  providedIn: 'root'
})
export class CheckStatusService {

  constructor(private http: HttpClient) { }

  // getWebsiteStatusList(websites: string[]): Observable<WebsiteCheckerResponse> {
  //   const request = {'websites': websites}
  //   return this.http.post<WebsiteCheckerResponse>("/api/v1/webchecker/websites", request);
  // }

  uploadCSV(f: FormData): Observable<any | WebsiteCheckerResponse> {
    return this.http.post<any | WebsiteCheckerResponse>("/api/v1/webchecker/websites", f, {reportProgress: true, observe: 'events'});
  }
}
