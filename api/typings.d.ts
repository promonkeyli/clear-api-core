declare namespace API {
    type Tool = {
           name?:string;
           url?:string;
           category?:string;
    }
    type User = {
           password?:string;
           username?:string;
    }
    type Response = {
           /** 响应时间戳 */
           time?:number;
           /** 响应码 */
           code?:StatusCode;
           /** 响应数据(可以为空) */
           data?:any;
           /** 响应描述 */
           message?:string;
    }
   type StatusCode = 200 | 400 | 401 | 403 | 404 | 500 | 501 | 502 | 503;
}