declare namespace API {
    type ApiResponse = {
           code?:number;
           message?:string;
           type?:string;
    }
    type Category = {
           name?:string;
           id?:number;
    }
    type Order = {
           complete?:boolean;
           id?:number;
           petId?:number;
           quantity?:number;
           shipDate?:string;
           /** Order Status */
           status?:string;
    }
    type Pet = {
           category?:Category;
           id?:number;
           name:string;
           photoUrls:Array<string>;
           /** pet status in the store */
           status?:string;
           tags?:Array<Tag>;
    }
    type Tag = {
           id?:number;
           name?:string;
    }
    type User = {
           lastName?:string;
           password?:string;
           phone?:string;
           /** User Status */
           userStatus?:number;
           username?:string;
           email?:string;
           firstName?:string;
           id?:number;
    }
}