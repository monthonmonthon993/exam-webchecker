import { HttpEventType } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { finalize, Subscription } from 'rxjs';
import { CheckStatusService } from '../check-status.service';
import { WebsiteCheckerResponse } from '../website-status';

@Component({
  selector: 'app-websites-checker',
  templateUrl: './websites-checker.component.html',
  styleUrls: ['./websites-checker.component.scss']
})
export class WebsitesCheckerComponent implements OnInit {
  uploadProgressNum: number = 0;
  processProgressNum: number = 0;
  processing = false;
  uploading = false;
  uploadSub: Subscription = new Subscription();
  resultReceived = false;
  upCount = 0;
  downCount = 0;
  timeCount = { minutes: 0, seconds: 0 };
  filename = '';

  constructor(private uploadService: CheckStatusService) { }

  ngOnInit(): void {
  }

  onFileSelected(event: any) {
    const file: File = event.target.files[0];
    this.readCSV(file)
  }

  onDragAndDropUpload(fs: FileList) {
    this.readCSV(fs[0])
  }

  private readCSV(file: File) {
    const startDate = new Date();
    this.uploadProgressNum = 0;
    this.processProgressNum = 0;
    this.resultReceived = false;

    if (file != null) {
      this.uploading = true;
      this.filename = file.name;
    } else {
      this.uploading = false;
      this.processing = false;
      return;
    }

    const formData = new FormData();

    formData.append("csv", file);

    const upload$ = this.uploadService.uploadCSV(formData)
      .pipe(
        finalize(() => {
          this.uploadSub.unsubscribe();
        })
      )

    this.uploadSub = upload$.subscribe({
      next: event => {
        if (event.type == HttpEventType.Sent) {
          // sent 
          console.log('start upload')
        }
        if (event.type == HttpEventType.UploadProgress) {
          console.log('uploading')
          this.uploadProgressNum = Math.round(100 * (event.loaded / event.total));
          this.uploading = false;
          this.processing = true;
        }
        if (event.type == HttpEventType.ResponseHeader) {
          console.log('response header received!');
        }
        if (event.type == HttpEventType.Response) {
          const res: WebsiteCheckerResponse = event.body;
          if (res.website_status_list && res.website_status_list.length > 0) {
            this.upCount = res.website_status_list.filter(el => el.ok === true).length;
            this.downCount = res.website_status_list.filter(el => el.ok === false).length;
            this.processProgressNum = 100;
            this.resultReceived = true;
            const duration = new Date().getTime() - startDate.getTime();
            const d = new Date(duration)
            this.timeCount = { minutes: d.getMinutes(), seconds: d.getSeconds() + d.getMilliseconds() / 1000 }
            this.processing = false;
          }
        }
      },
      error: (e) => {
        this.uploadProgressNum = 100;
        this.processProgressNum = 0;
        this.uploading = false;
        this.processing = false;
        console.error(e)
      },
      complete: () => {
        
      }
    });
  }
}
