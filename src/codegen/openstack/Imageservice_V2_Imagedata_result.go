package openstack

// Code generated by cloud manager.

import (
)


// call result's extract function
func ExtractUploadImageserviceV2ImagedataResponse(response *UploadImageserviceV2ImagedataResponse)(interface{}, error){
    return response.UploadResult.Body, response.UploadResult.Err
}



// call result's extract function
func ExtractStageImageserviceV2ImagedataResponse(response *StageImageserviceV2ImagedataResponse)(interface{}, error){
    return response.StageResult.Body, response.StageResult.Err
}



// call result's extract function
func ExtractDownloadImageserviceV2ImagedataResponse(response *DownloadImageserviceV2ImagedataResponse)(interface{}, error){
    return response.DownloadResult.Body, response.DownloadResult.Err
}
