import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { WebsitesCheckerComponent } from './websites-checker/websites-checker.component';

const routes: Routes = [
  { path: '', redirectTo: '/websites-checker', pathMatch: 'full'},
  { path: 'websites-checker', component: WebsitesCheckerComponent},
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
