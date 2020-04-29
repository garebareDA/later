import axios from "axios";
function post(url:string, params:object):boolean {
  try{
    axios.post(url, params);
    return false;
  }catch{
    return true;
  }
}

export default{
  post
}