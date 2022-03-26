export interface WebsiteStatus {
    name: string;
    ok: boolean;
    err_msg: string;
}

export interface WebsiteCheckerResponse {
    website_status_list: WebsiteStatus[];
}
