declare namespace API {
    type Pet = {
           id?:number;
           name:string;
           photoUrls:Array<string>;
           /** pet status in the store */
           status?:string;
           tags?:Array<Tag>;
           category?:Category;
    }
    type Tag = {
           id?:number;
           name?:string;
    }
    type User = {
           username?:string;
           email?:string;
           firstName?:string;
           id?:number;
           lastName?:string;
           password?:string;
           phone?:string;
           /** User Status */
           userStatus?:number;
    }
    type ApiResponse = {
           message?:string;
           type?:string;
           code?:number;
    }
    type Category = {
           name?:string;
           id?:number;
    }
    type Order = {
           /** Order Status */
           status?:string;
           complete?:boolean;
           id?:number;
           petId?:number;
           quantity?:number;
           shipDate?:string;
    }
}