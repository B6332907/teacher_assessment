import React from "react";
import { Teacher_assessmentsInterface } from "../models/ITeacher_assessment";

const apiUrl = "http://localhost:8080";

async function GetStudent() {
    const requestOptions = {
        method: "GET",
        headers: {
            //Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json"
        },
    };

    let res = await fetch(`${apiUrl}/Students`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data);
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

async function GetTeacher() {
    const requestOptions = {
        method: "GET",
        headers: {
            // Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json"
        },
    };

    let res = await fetch(`${apiUrl}/teachers`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data);
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

async function GetTeaching_duration() {
    const requestOptions = {
        method: "GET",
        headers: {
            //Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json"
        },
    };

    let res = await fetch(`${apiUrl}/Teaching_durations`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data);
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

async function GetContent_difficulty_level() {
    const requestOptions = {
        method: "GET",
        headers: {
            //Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json"
        },
    };

    let res = await fetch(`${apiUrl}/Content_difficulty_levels`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data);
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

async function GetContent_quality() {
    const requestOptions = {
        method: "GET",
        headers: {
            //Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json"
        },
    };

    let res = await fetch(`${apiUrl}/Content_qualitys`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data);
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

async function GetTeacher_assessments() {
    const requestOptions = {
        method: "GET",
        headers: {
            //Authorization: `Bearer ${localStorage.getItem("token")}`,
            "Content-Type": "application/json"
        },
    };

    let res = await fetch(`${apiUrl}/Teacher_assessments`, requestOptions)
        .then((response) => response.json())
        .then((res) => {
            if (res.data) {
                console.log(res.data);
                return res.data;
            } else {
                return false;
            }
        });

    return res;
}

async function CreateTeacher_assessment(data: Teacher_assessmentsInterface) {
    const requestOptions = {
      method: "POST",
      headers: {
        //Authorization: `Bearer ${localStorage.getItem("token")}`,
        "Content-Type": "application/json",
      },
      body: JSON.stringify(data),
    };
  
    let res = await fetch(`${apiUrl}/Teacher_assessments`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        if (res.data) {
          return res.data;
        } else {
          return false;
        }
      });
  
    return res;
  }

  export {
    GetTeacher_assessments,
    GetStudent,
    GetTeacher,
    GetTeaching_duration,
    GetContent_difficulty_level,
    GetContent_quality,
    CreateTeacher_assessment,
  };






