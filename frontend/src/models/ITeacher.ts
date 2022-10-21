import { StudentsInterface } from "./IStudent";

export interface TeachersInterface {

    ID?: number,
    Faculty_id?: string;
    Level?: string;
    Name?: string;
    Email?: string;
    Graduate_faculty_level_id?: number;
    Officer_id?: number;

    Student?: StudentsInterface
   }
