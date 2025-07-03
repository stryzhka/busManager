import {Button, MenuListItem, ScrollView, styleReset, TextInput} from "react95";
import React, {useEffect, useState} from "react";
import {createGlobalStyle} from "styled-components";
import ms_sans_serif from "../assets/fonts/fixedsys.woff2";
import ms_sans_serif_bold from 'react95/dist/fonts/ms_sans_serif_bold.woff2';
import original from 'react95/dist/themes/original';
import {Add, GetAll} from "../../wailsjs/go/routers/BusRouter.js";
import {GetById} from "../../wailsjs/go/routers/BusRouter.js";

const GlobalStyles = createGlobalStyle`
  ${styleReset}
  @font-face {
    font-family: 'ms_sans_serif';
    src: url('${ms_sans_serif}') format('woff2');
    font-weight: 400;
    font-style: normal
  }
  @font-face {
    font-family: 'ms_sans_serif';
    src: url('${ms_sans_serif_bold}') format('woff2');
    font-weight: bold;
    font-style: normal
  }
  body, input, select, textarea {
    font-family: 'ms_sans_serif';
  }
`;



const BusComponent = () => {
    const [items, setItems] = useState([]);
    const [selectedItem, setSelectedItem] = useState(null);
    useEffect(() => {
        GetAll().then(
            result => {

                console.log(JSON.parse(result))
                setItems(JSON.parse(result))
            }
        )
    }, []);
    const handleItemClick = (item) => {
        console.log("ХУЙ")

        GetById(item.ID).then(
            result => {
                const selectedData = JSON.parse(result);
                setSelectedItem(selectedData);
                console.log(selectedData)
            }
        ).catch(err => console.error("Ошибка при получении данных по ID:", err));
    };
    const handleInputChange = (field) => (e) => {
        setSelectedItem(prev => ({
            ...prev,
            [field]: e.target.value
        }));
    };
    const handleCreate = () => {
        if (selectedItem) {
            selectedItem.ID = null
            Add(JSON.stringify(selectedItem)).then(
                result => {
                    if (JSON.parse(result).Error){
                        console.log(JSON.parse(result).Error)
                    }

                    GetAll().then(
                        result => {
                            setItems(JSON.parse(result));
                            setSelectedItem(null);
                        }
                    ).catch(err => console.error("Ошибка при обновлении списка:", err));
                }
            ).catch(err => console.error("Ошибка при создании:", err));
        } else {
            console.warn("Нет выбранного элемента для создания");
        }
    };
    return (
        <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start' }}>
            <ScrollView style={{width: '200px', height: '200px', marginRight: '20px'}}>
                {items.map((item, index) => (
                    <MenuListItem key={index} onClick={() => handleItemClick(item)}>
                        {item.Brand}
                    </MenuListItem>
                ))}

            </ScrollView>
            <div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}} value={selectedItem?.Brand || ''}
                               onChange={handleInputChange('Brand')}></TextInput>
                    <div style={{marginRight: '20px'}}>Марка</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}} value={selectedItem?.BusModel || ''}
                               onChange={handleInputChange('BusModel')}></TextInput>
                    <div style={{marginRight: '20px'}}>Модель</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}} value={selectedItem?.RegisterNumber || ''}
                        onChange={handleInputChange('RegisterNumber')}></TextInput>
                    <div style={{marginRight: '20px'}}>Регистрационный номер</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}} value={selectedItem?.AssemblyDate || ''}
                               onChange={handleInputChange('AssemblyDate')}></TextInput>
                    <div style={{marginRight: '20px'}}>Дата смерти</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start', marginBottom: '10px' }}>
                    <TextInput style={{width: '150px', marginRight: '20px'}} value={selectedItem?.LastRepairDate || ''}
                               onChange={handleInputChange('LastRepairDate')}></TextInput>
                    <div style={{marginRight: '20px'}}>Дата гниения</div>
                </div>
                <div style={{ display: 'flex', flexDirection: 'row', alignItems: 'flex-start',  }}>
                    <Button style={{marginRight: '10px'}}>Сохранить</Button>
                    <Button style={{marginRight: '10px'}}>Удалить</Button>
                    <Button onClick={handleCreate}>Создать</Button>
                </div>
            </div>

        </div>
    )
}

export default BusComponent